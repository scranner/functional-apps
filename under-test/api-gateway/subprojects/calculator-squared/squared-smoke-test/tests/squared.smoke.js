const { API_UNDER_TEST /* , PUSH_RESULTS */ } = process.env;

const axios = require('axios');

const timeFunction = async (functionToTime, url, params) => {
  const t0 = Date.now();
  let result;
  try {
    result = await functionToTime(url, params);
  } catch (e) {
    result = e;
  }
  const t1 = Date.now();
  return {
    result,
    timing: t1 - t0,
  };
};


describe('When a number of requests are made to the service', () => {
  const timings = [];

  describe('Which are valid', () => {
    it('2^2 = 4', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 2 } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('4');
    });

    it('"10"^2 = 100', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '10' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('100');
    });

    it('"10e2"^2 = 1000000', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '10e2' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('1000000');
    });

    describe('Which are invalid', () => {
      it('"test"^2 = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 'test' } });
        timings.push(result.timing);
        expect(result.result.response.status).toBe(400);
      });

      it('noting passed = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`);
        timings.push(result.timing);
        expect(result.result.response.status).toBe(400);
      });

      afterAll(() => {
        // eslint-disable-next-line no-console
        // if (PUSH_RESULTS === 'true') {
        //   axios.post('monitoring service url', {
        //     timings: timings.reduce((a, b) => a + b, 0) / timings.length,
        //   });
        // }
        console.log(`Average Request timing: ${timings.reduce((a, b) => a + b, 0) / timings.length}`);
      });
    });
  });
});
