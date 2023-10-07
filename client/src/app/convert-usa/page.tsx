"use client";

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
import { useState } from "react";
import { getRequest, postRequest } from "../../api/fetchRequests";

export default function Page() {
  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: "usa",
    inputUnit: "",
    outputSystem: "metric",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<singleOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });
  const inputComplete = Object.values(input).every((item) => !!item);

  const volumeInputs = ["cups", "gallons", "quarts", "pints", "fluid oz"];
  const weightInputs = ["cups", "lbs", "oz"];
  const volumeOutputs = ["millilitres", "litres"];
  const weightOutputs = ["grams", "kg"];

  const handleInput = <K extends keyof singleInput>(
    property: K,
    value: singleInput[K]
  ) => {
    setInput({
      ...input,
      [property]: value === "string" ? value.toLowerCase() : value,
    });
  };

  const getSampleData = async () => {
    const data = await getRequest();
    setOutput(data);
  };

  const convertVolumeUS = async () => {
    const data = await postRequest(
      "http://localhost:8080/api/convert/volume-us",
      input
    );
    setOutput(data);
  };

  const convertWeightUS = async () => {
    const data = await postRequest(
      "http://localhost:8080/api/convert/weight-us",
      input
    );
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
        onClick={input.type === "volume" ? convertVolumeUS : convertWeightUS}
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
