import { HeightFeet, HeightMetric, ValidInputs } from "@/types/heightTypes";

export const inputComplete = (
  input: ValidInputs,
  heightFeet: HeightFeet,
  heightMetric: HeightMetric
): boolean => {
  switch (input) {
    case "feet":
      return heightFeet.feet >= 0 || heightFeet.inches >= 0;
    case "centimetres":
      return heightMetric.centimetres >= 0;
    case "metres":
      return heightMetric.metres >= 0;
  }
};
