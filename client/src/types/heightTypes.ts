export type HeightFeet = {
  feet: number;
  inches: number;
};

export type HeightMetric = {
  centimetres: number;
  metres: number;
};

const INPUTS = ["centimetres", "metres", "feet"] as const;
type InputsTuple = typeof INPUTS;
export type ValidInputs = InputsTuple[number];

export type HeightOutput = (HeightFeet | HeightMetric) & {
  error?: string;
};
