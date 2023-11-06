import { postRequest } from "@/api/fetchRequests";
import { Button } from "@/components/ui/button";
import {
  ConversionSystem,
  ConversionType,
  SingleInput,
  SingleOutput,
} from "@/types/conversionTypes";
import { inputComplete } from "@/utils/recipe";
import { useState } from "react";
import MultipleInputsComp from "./MultipleInputs";
import AddInputDropdown from "./AddInputDropdown";

export default function MultipleConversions() {
  const [conversionSystem, setConversionSystem] =
    useState<ConversionSystem>("usa");
  const [conversionType, setConversionType] =
    useState<ConversionType>("volume");

  const [input, setInput] = useState<SingleInput>({
    ingredient: "",
    inputSystem: conversionSystem === "usa" ? "usa" : "metric",
    inputUnit: "",
    outputSystem: conversionSystem === "usa" ? "usa" : "metric",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<SingleOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });
  const [inputList, setInputList] = useState<Array<SingleInput>>([]);
  const [outputList, setOutputList] = useState<Array<SingleOutput>>([]);

  const inputListComplete =
    !inputList.length ||
    (!!inputList.length && !!inputList.every((input) => inputComplete(input)));

  const handleListConversion = async () => {
    let data = [{ ingredient: "", unit: "", amount: 0 }];
    data = await postRequest("list", inputList);
    setOutputList(data);
  };

  return (
    <div className="text-center">
      <MultipleInputsComp inputList={inputList} setInputList={setInputList} />
      <AddInputDropdown inputList={inputList} setInputList={setInputList} />

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
