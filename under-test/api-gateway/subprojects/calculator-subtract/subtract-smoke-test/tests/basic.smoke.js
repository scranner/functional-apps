const { API_UNDER_TEST } = process.env;

const axios = require('axios');

describe('When a request is made to the service', () => {
  describe('Which is valid', () => {
    let result;

    beforeAll(async () => {
      result = await axios.get(`http://${API_UNDER_TEST}`, { params: { x: 2, y: 4 } });
    });

    it('should return a 200', () => {
      expect(result.status).toBe(200);
    });
    it('should return the result', () => {
      expect(result.data.result).toBeDefined();
    });
  });

  describe('Which is invalid', () => {
    let result;

    beforeAll(async () => {
      try {
        result = await axios.get(`http://${API_UNDER_TEST}`);
      } catch (e) {
        result = e;
      }
    });

    it('should return a 400', () => {
      expect(result.response.status).toBe(400);
    });
  });
});
