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
    it('1 + 2 = 3', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 1, y: 2 } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('3');
    });

    it('"10" + "5" = 15', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '10', y: '5' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('15');
    });

    it('"10e2" + "12" = 1012', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '10e2', y: '12' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('1012');
    });

    describe('Which are invalid', () => {
      it('"test" + "5" = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 'test', y: '5' } });
        timings.push(result.timing);
        expect(result.result.response.status).toBe(400);
      });

      it('"10" + "wee" = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '10', y: 'wee' } });
        timings.push(result.timing);
        expect(result.result.response.status).toBe(400);
      });

      it('noting passed + "beep" = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { y: 'beep' } });
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
