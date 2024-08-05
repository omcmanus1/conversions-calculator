export type HeightFeet = {
  feet: number;
  inches: number;
};

export type HeightMetric = {
  centimetres: number;
  metres: number;
};

export enum ValidInputs {
  centimetres = "Centimetres",
  metres = "Metres",
  feet = "Feet",
}
