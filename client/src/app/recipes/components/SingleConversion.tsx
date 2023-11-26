import { postRequest } from "@/api/fetchRequests";
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { ConversionSystem, RecipeInput, RecipeOutput } from "@/types/conversionTypes";
import { inputComplete } from "@/utils/recipe";
import { useState } from "react";
import SingleInputComp from "./SingleInput";

export type Props = {
  conversionType: ConversionSystem;
};

export default function SingleConversion({ conversionType }: Props) {
  const [input, setInput] = useState<RecipeInput>({
    ingredient: "",
    inputSystem: conversionType === "usa" ? "usa" : "metric",
    inputUnit: "",
    outputSystem: conversionType === "usa" ? "metric" : "usa",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<RecipeOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });

  const handleSingleConversion = async () => {
    let data: RecipeOutput;
    switch (input.inputSystem) {
      case "usa":
        input.type === "volume"
          ? (data = await postRequest("volume-us", input))
          : (data = await postRequest("weight-us", input));
        break;
      case "metric":
        input.type === "volume"
          ? (data = await postRequest("volume-metric", input))
          : (data = await postRequest("weight-metric", input));
        break;
    }
    setOutput(data);
  };

  return (
    <div className="text-center">
      <SingleInputComp
        input={input}
        setInput={setInput}
        conversionType={conversionType}
      />
      <Button
        className={`mt-3 mb-3 ${inputComplete(input) && "hover:bg-lime-100"}`}
        disabled={!inputComplete(input)}
        variant="outline"
        onClick={handleSingleConversion}
      >
        {input.inputSystem === "usa" ? "Freedom" : "Metric"}
        <ChevronDoubleRight className="w-5" />
        {input.inputSystem === "usa" ? "Metric" : "Freedom"}
      </Button>
      {!!output?.amount && (
        <Card>
          <CardHeader>
            <CardTitle>{output?.ingredient}</CardTitle>
            <CardDescription>{`${output?.amount} ${output?.unit}`}</CardDescription>
          </CardHeader>
        </Card>
      )}
    </div>
  );
}
