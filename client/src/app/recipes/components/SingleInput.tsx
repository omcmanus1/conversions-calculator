import SelectSh from "@/components/select";
import { Input } from "@/components/ui/input";
import {
  METRIC_VOLUME,
  METRIC_WEIGHT,
  US_VOLUME,
  US_WEIGHT,
} from "@/constants/measures";
import { conversionTypes, singleInput } from "@/types/conversionTypes";
import { Dispatch, SetStateAction } from "react";
import { handleInput } from "@/utils/recipe";

interface Props {
  input: singleInput;
  setInput: Dispatch<SetStateAction<singleInput>>;
  conversionType: conversionTypes;
}

export default function SingleInput({
  input,
  setInput,
  conversionType,
}: Props) {
  const weightInputs = conversionType === "usa" ? US_WEIGHT : METRIC_WEIGHT;
  const weightOutputs = conversionType === "usa" ? METRIC_WEIGHT : US_WEIGHT;
  const volumeInputs = conversionType === "usa" ? US_VOLUME : METRIC_VOLUME;
  const volumeOutputs = conversionType === "usa" ? METRIC_VOLUME : US_VOLUME;
  return (
    <>
      <Input
        className="mb-1"
        type="text"
        placeholder="Ingredient"
        onChange={(e) =>
          handleInput(setInput, input, "ingredient", e.target.value)
        }
      />
      <SelectSh
        disabled={!input.ingredient}
        handleChange={(e) => handleInput(setInput, input, "type", e)}
        placeholder="Type"
        selectContent={["weight", "volume"]}
      />
      <SelectSh
        disabled={!input.type}
        handleChange={(e) => handleInput(setInput, input, "inputUnit", e)}
        placeholder="Input Unit"
        selectContent={
          input.type === "weight"
            ? weightInputs
            : input.type === "volume"
            ? volumeInputs
            : ["Please specify a unit"]
        }
      />
      <Input
        className="mb-1"
        disabled={!input.inputUnit}
        type="number"
        placeholder="Amount"
        onChange={(e) =>
          handleInput(setInput, input, "amount", Number(e.target.value))
        }
      />
      <SelectSh
        disabled={!input.amount}
        handleChange={(e) => handleInput(setInput, input, "outputUnit", e)}
        placeholder="Output Unit"
        selectContent={
          input.type === "weight"
            ? weightOutputs
            : input.type === "volume"
            ? volumeOutputs
            : ["Please specify a unit"]
        }
      />
    </>
  );
}
