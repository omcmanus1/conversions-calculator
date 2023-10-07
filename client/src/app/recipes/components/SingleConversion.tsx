import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import SelectSh from "@/components/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { singleInput, singleOutput } from "@/types/conversionTypes";
import React from "react";
import { postRequest } from "@/api/fetchRequests";

export default function SingleConversion({
  input,
  setInput,
  output,
  setOutput,
  weightInputs,
  volumeInputs,
  weightOutputs,
  volumeOutputs,
  list = false,
}: {
  input: singleInput;
  setInput: React.Dispatch<React.SetStateAction<singleInput>>;
  output: singleOutput;
  setOutput: React.Dispatch<React.SetStateAction<singleOutput>>;
  weightInputs: Array<string>;
  volumeInputs: Array<string>;
  weightOutputs: Array<string>;
  volumeOutputs: Array<string>;
  list?: boolean;
}) {
  const inputComplete = Object.values(input).every((item) => !!item);

  const handleInput = <K extends keyof singleInput>(
    property: K,
    value: singleInput[K]
  ) => {
    setInput({
      ...input,
      [property]: value === "string" ? value.toLowerCase() : value,
    });
  };

  const handleConversion = async () => {
    let data = { ingredient: "", unit: "", amount: 0 };
    input.type === "volume"
      ? (data = await postRequest("volume-us", input))
      : (data = await postRequest("weight-us", input));
    setOutput(data);
  };

  return (
    <div className="text-center">
      <Input
        className="mb-1"
        type="text"
        placeholder="Ingredient"
        onChange={(e) => handleInput("ingredient", e.target.value)}
      />
      <SelectSh
        handleChange={(e) => handleInput("type", e)}
        placeholder="Type"
        selectContent={["weight", "volume"]}
      />
      <SelectSh
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
        type="number"
        placeholder="Amount"
        onChange={(e) => handleInput("amount", Number(e.target.value))}
      />
      <SelectSh
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
      <Button
        className={`mt-3 mb-3 ${inputComplete && "hover:bg-lime-100"}`}
        disabled={!inputComplete}
        variant="outline"
        onClick={handleConversion}
      >
        Freedom
        <ChevronDoubleRight className="w-5" />
        Metric
      </Button>
      {!!output?.amount && (
        <Card>
          <CardHeader>
            <CardTitle>{output?.ingredient}</CardTitle>
            <CardDescription>
              {`${output?.amount} ${output?.unit}`}
            </CardDescription>
          </CardHeader>
        </Card>
      )}
    </div>
  );
}
