import SelectSh from "@/components/select";
import { Input } from "@/components/ui/input";
import { singleInput } from "@/types/conversionTypes";

interface Props {
  input: singleInput;
  weightInputs: Array<string>;
  volumeInputs: Array<string>;
  weightOutputs: Array<string>;
  volumeOutputs: Array<string>;
  handleInput: <K extends keyof singleInput>(
    property: K,
    value: singleInput[K]
  ) => void;
}

export default function SingleInput({
  input,
  weightInputs,
  volumeInputs,
  weightOutputs,
  volumeOutputs,
  handleInput,
}: Props) {
  return (
    <>
      <Input
        className="mb-1"
        type="text"
        placeholder="Ingredient"
        onChange={(e) => handleInput("ingredient", e.target.value)}
      />
      <SelectSh
        disabled={!input.ingredient}
        handleChange={(e) => handleInput("type", e)}
        placeholder="Type"
        selectContent={["weight", "volume"]}
      />
      <SelectSh
        disabled={!input.type}
        handleChange={(e) => handleInput("inputUnit", e)}
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
        onChange={(e) => handleInput("amount", Number(e.target.value))}
      />
      <SelectSh
        disabled={!input.amount}
        handleChange={(e) => handleInput("outputUnit", e)}
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
