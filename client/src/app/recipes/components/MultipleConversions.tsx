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
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";

export default function MultipleConversions() {
  const [inputList, setInputList] = useState<Array<SingleInput>>([]);
  const [outputList, setOutputList] = useState<Array<SingleOutput>>([]);

  const inputListComplete =
    !!inputList.length && !!inputList.every((input) => inputComplete(input));

  const handleListConversion = async () => {
    let data: SingleOutput[];
    data = await postRequest("list", inputList);
    setOutputList(data);
  };

  return (
    <div className="text-center">
      <div className="text-center">
        <MultipleInputsComp inputList={inputList} setInputList={setInputList} />
        <AddInputDropdown inputList={inputList} setInputList={setInputList} />
      </div>
      <Button
        className={`mt-3 mb-3 ${!inputListComplete && "hover:bg-lime-100"}`}
        disabled={!inputListComplete}
        variant="outline"
        onClick={handleListConversion}
      >
        Convert All
        <ChevronDoubleRight className="w-5" />
      </Button>
      {/* {!!output?.amount && (
        <Card>
        <CardHeader>
        <CardTitle>{output?.ingredient}</CardTitle>
        <CardDescription>{`${output?.amount} ${output?.unit}`}</CardDescription>
        </CardHeader>
        </Card>
      )} */}
    </div>
  );
}
