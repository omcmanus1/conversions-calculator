export type conversionTypes = "usa" | "metric";

export type singleInput = {
  ingredient: string;
  inputSystem: conversionTypes;
  inputUnit: string;
  outputSystem: conversionTypes;
  outputUnit: string;
  type: string;
  amount: number;
};

export type singleOutput = {
  ingredient: string;
  unit: string;
  amount: number;
};

export type usVolume = ["cups", "gallons", "quarts", "pints", "fluid oz"];
export type usWeight = ["cups", "lbs", "oz"];
export type metricVolume = ["millilitres", "litres"];
export type metricWeight = ["grams", "kg"];

export type measures = {
  usVolume: Array<usVolume>;
  usWeight: Array<usWeight>;
  metricVolume: Array<metricVolume>;
  metricWeight: Array<usVolume>;
};
