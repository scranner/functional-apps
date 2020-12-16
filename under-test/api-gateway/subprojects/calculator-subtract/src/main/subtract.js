const getInt = x => {
  const parsed = parseInt(x);
  if (isNaN(parsed)) return undefined;
  return parsed;
};

const subtract = (x, y) => {
    console.log(`Processing resquest for ${x} and ${y}`);
    const x1 = Number(x);
    const y1 = Number(y);

    if (!x1 || !y1) {
        return undefined;
    }

    return x1 - y1;
};

module.exports = { subtract };