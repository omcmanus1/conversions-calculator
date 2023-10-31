import { postRequest } from "@/api/fetchRequests";
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  conversionTypes,
  singleInput,
  singleOutput,
} from "@/types/conversionTypes";
import { useState } from "react";
import SingleInput from "./SingleInput";
import { PlusIcon } from "lucide-react";
import { inputComplete } from "@/utils/recipe";

export default function MultipleConversions() {
  const [conversionType, setConversionType] = useState<conversionTypes>("usa");
  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: conversionType === "usa" ? "usa" : "metric",
    inputUnit: "",
    outputSystem: conversionType === "usa" ? "usa" : "metric",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<singleOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });
  const [inputList, setInputList] = useState<Array<singleInput>>([]);
  const [outputList, setOutputList] = useState<Array<singleOutput>>([]);

  const inputListComplete =
    !inputList.length ||
    (!!inputList.length && !!inputList.every((input) => inputComplete(input)));

  const addInputBox = (convType: conversionTypes) => {
    setConversionType(convType);
    setInputList([...inputList, input]);
  };

  const handleListConversion = async () => {
    let data = [{ ingredient: "", unit: "", amount: 0 }];
    data = await postRequest("list", inputList);
    setOutputList(data);
  };

  return (
    <div className="text-center">
      {!!inputList.length &&
        inputList.map((currentInput, index) => {
          return (
            <SingleInput
              key={`input_${index}`}
              input={input}
              setInput={setInput}
              conversionType={currentInput.inputSystem}
            />
          );
        })}
      <Button
        className={`mt-3 mb-3 ${!!inputListComplete && "hover:bg-teal-100"}`}
        disabled={!inputListComplete}
        variant="outline"
        onClick={() => addInputBox("usa")}
      >
        Add US Conversion
      </Button>
      <Button
        className={`mb-1 ${!!inputListComplete && "hover:bg-teal-100"}`}
        disabled={!inputListComplete}
        variant="outline"
        onClick={() => addInputBox("metric")}
      >
        Add Metric Conversion
      </Button>

      {/* <Button
        className={`mt-3 mb-3 ${!!inputList.length && "hover:bg-lime-100"}`}
        disabled={!inputList.length}
        variant="outline"
        onClick={handleListConversion}
      >
        Convert All
        <ChevronDoubleRight className="w-5" />
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
      )} */}
    </div>
  );
}
