import { Dispatch, SetStateAction } from "react";

export type ConversionSystem = "usa" | "metric";

export type ConversionType = "volume" | "weight";

export type InputFields = "ingredient" | "amount" | "inputUnit" | "outputUnit";

export type RecipeInput = {
  ingredient: string;
  inputSystem: ConversionSystem;
  inputUnit: string;
  outputSystem: ConversionSystem;
  outputUnit: string;
  type: string;
  amount: number;
};

export type RecipeOutput = {
  ingredient: string;
  unit: string;
  amount: number;
};

export type InputListProps = {
  inputList: RecipeInput[];
  setInputList: Dispatch<SetStateAction<RecipeInput[]>>;
};
