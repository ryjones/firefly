import { setup, shutDown, app, mockEventStreamWebSocket } from './common';
import nock from 'nock';
import request from 'supertest';
import assert from 'assert';
import { IEventMemberRegistered } from '../lib/interfaces';

describe('Members', async () => {

  before(async () => {
    await setup();
  });

  after(() => {
    shutDown();
  });

  it('Checks that an empty array is initially returned when querying members', async () => {
    const result = await request(app)
      .get('/api/v1/members')
      .expect(200);
    assert.deepStrictEqual(result.body, []);
  });

  it('Attempting to add a member without an address should raise an error', async () => {
    const result = await request(app)
      .put('/api/v1/members')
      .send({
        name: 'Member A',
        app2appDestination: 'kld://app2app',
        docExchangeDestination: 'kld://docexchange'
      })
      .expect(400);
    assert.deepStrictEqual(result.body, { error: 'Invalid member' });
  });

  it('Attempting to add a member without a name should raise an error', async () => {
    const result = await request(app)
      .put('/api/v1/members')
      .send({
        address: '0x0000000000000000000000000000000000000001',
        app2appDestination: 'kld://app2app',
        docExchangeDestination: 'kld://docexchange'
      })
      .expect(400);
    assert.deepStrictEqual(result.body, { error: 'Invalid member' });
  });

  it('Attempting to add a member without an app2app destination should raise an error', async () => {
    const result = await request(app)
      .put('/api/v1/members')
      .send({
        address: '0x0000000000000000000000000000000000000001',
        name: 'Member A',
        docExchangeDestination: 'kld://docexchange'
      })
      .expect(400);
    assert.deepStrictEqual(result.body, { error: 'Invalid member' });
  });

  it('Attempting to add a member without a document exchange destination should raise an error', async () => {
    const result = await request(app)
      .put('/api/v1/members')
      .send({
        address: '0x0000000000000000000000000000000000000001',
        name: 'Member A',
        app2appDestination: 'kld://app2app',
      })
      .expect(400);
    assert.deepStrictEqual(result.body, { error: 'Invalid member' });
  });

  it('Checks that adding a member sends a request to API Gateway and updates the database', async () => {

    nock('https://apigateway.kaleido.io')
      .post('/registerMember?kld-from=0x0000000000000000000000000000000000000001&kld-sync=false')
      .reply(200);
    const addMemberResponse = await request(app)
      .put('/api/v1/members')
      .send({
        address: '0x0000000000000000000000000000000000000001',
        name: 'Member A',
        app2appDestination: 'kld://app2app',
        docExchangeDestination: 'kld://docexchange'
      })
      .expect(200);
    assert.deepStrictEqual(addMemberResponse.body, { status: 'submitted' });

    const getMemberResponse = await request(app)
      .get('/api/v1/members')
      .expect(200);
    const member = getMemberResponse.body[0];
    assert.strictEqual(member.address, '0x0000000000000000000000000000000000000001');
    assert.strictEqual(member.name, 'Member A');
    assert.strictEqual(member.app2appDestination, 'kld://app2app');
    assert.strictEqual(member.docExchangeDestination, 'kld://docexchange');
    assert.strictEqual(member.owned, true);
    assert.strictEqual(member.confirmed, false);
    assert.strictEqual(typeof member.timestamp, 'number');
  });

  it('Checks that event stream notification for confirming member registrations are handled', async () => {

    const eventPromise = new Promise((resolve) => {
      mockEventStreamWebSocket.once('send', message => {
        assert.strictEqual(message, '{"type":"ack","topic":"dev"}');
        resolve();
      })
    });

    const data: IEventMemberRegistered = {
      member: '0x0000000000000000000000000000000000000001',
      name: 'Member A',
      app2appDestination: 'kld://app2app',
      docExchangeDestination: 'kld://docexchange',
      timestamp: 1
    };
    mockEventStreamWebSocket.emit('message', JSON.stringify([{
      signature: 'MemberRegistered(address,string,string,string,uint256)',
      data
    }]));
    await eventPromise;
  });

  it('Get member should return the confirmed member', async () => {
    const result = await request(app)
      .get('/api/v1/members')
      .expect(200);
    const member = result.body[0];
    assert.strictEqual(member.confirmed, true);
  });

  it('Checks that updating a member sends a request to API Gateway and updates the database', async () => {
    nock('https://apigateway.kaleido.io')
      .post('/registerMember?kld-from=0x0000000000000000000000000000000000000001&kld-sync=false')
      .reply(200);
    const addMemberResponse = await request(app)
      .put('/api/v1/members')
      .send({
        address: '0x0000000000000000000000000000000000000001',
        name: 'Member B',
        app2appDestination: 'kld://app2app2',
        docExchangeDestination: 'kld://docexchange2'
      })
      .expect(200);
    assert.deepStrictEqual(addMemberResponse.body, { status: 'submitted' });

    const getMemberResponse = await request(app)
      .get('/api/v1/members')
      .expect(200);
    const member = getMemberResponse.body[0];
    assert.strictEqual(member.address, '0x0000000000000000000000000000000000000001');
    assert.strictEqual(member.name, 'Member B');
    assert.strictEqual(member.app2appDestination, 'kld://app2app2');
    assert.strictEqual(member.docExchangeDestination, 'kld://docexchange2');
    assert.strictEqual(member.owned, true);
    assert.strictEqual(member.confirmed, false);
    assert.strictEqual(typeof member.timestamp, 'number');
  });

  it('Checks that event stream notification for confirming member registrations are handled', async () => {

    const eventPromise = new Promise((resolve) => {
      mockEventStreamWebSocket.once('send', message => {
        assert.strictEqual(message, '{"type":"ack","topic":"dev"}');
        resolve();
      })
    });

    const data: IEventMemberRegistered = {
      member: '0x0000000000000000000000000000000000000001',
      name: 'Member B',
      app2appDestination: 'kld://app2app2',
      docExchangeDestination: 'kld://docexchange2',
      timestamp: 1
    };
    mockEventStreamWebSocket.emit('message', JSON.stringify([{
      signature: 'MemberRegistered(address,string,string,string,uint256)',
      data
    }]));
    await eventPromise;
  });

  it('Get member should return the confirmed member', async () => {
    const result = await request(app)
      .get('/api/v1/members')
      .expect(200);
    const member = result.body[0];
    assert.strictEqual(member.confirmed, true);
  });

});
