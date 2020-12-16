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
    it('"4!" = 24', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '4' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('24');
    });

    it('3! = 6', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 3 } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('6');
    });

    it('"1e1"! = 3628800', async () => {
      const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: '1e1' } });
      timings.push(result.timing);
      expect(result.result.data.result).toBe('3628800');
    });

    describe('Which are invalid', () => {
      it('"mfwerf"! = Invalid Parameters, 400', async () => {
        const result = await timeFunction(axios.get, `http://${API_UNDER_TEST}`, { params: { x: 'mfwerf' } });
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
