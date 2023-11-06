export type ConversionSystem = "usa" | "metric";

export type ConversionType = "volume" | "weight";

export type InputFields = "ingredient" | "amount" | "inputUnit" | "outputUnit";

export type SingleInput = {
  ingredient: string;
  inputSystem: ConversionSystem;
  inputUnit: string;
  outputSystem: ConversionSystem;
  outputUnit: string;
  type: string;
  amount: number;
};

export type SingleOutput = {
  ingredient: string;
  unit: string;
  amount: number;
};
