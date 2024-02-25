import { postRequest } from "@/api/fetchRequests";
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { RecipeInput, RecipeOutput } from "@/types/conversionTypes";
import { inputComplete } from "@/utils/recipe";
import { useState } from "react";
import AddInputDropdown from "./AddInputDropdown";
import MultipleInputsComp from "./MultipleInputs";
import { LoadingSpinner } from "@/components/ui/loading-spinner";

export default function MultipleConversions() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [inputList, setInputList] = useState<RecipeInput[]>([]);
  const [outputList, setOutputList] = useState<RecipeOutput[]>([]);

  const inputListComplete =
    !!inputList.length && !!inputList.every((input) => inputComplete(input));

  const handleListConversion = async () => {
    setIsLoading(true);
    let data: RecipeOutput[];
    data = await postRequest("list", inputList);
    setOutputList(data);
    setIsLoading(false);
  };

  return (
    <div className="text-center">
      <div className="text-center">
        <MultipleInputsComp inputList={inputList} setInputList={setInputList} />
        <AddInputDropdown inputList={inputList} setInputList={setInputList} />
      </div>
      <Button
        className={`mt-3 mb-3 ${!inputListComplete && "hover:bg-lime-100"}`}
        disabled={!inputListComplete || isLoading}
        variant="outline"
        onClick={handleListConversion}
      >
        {isLoading ? (
          <LoadingSpinner />
        ) : (
          <>
            Convert All
            <ChevronDoubleRight className="w-5" />
          </>
        )}
      </Button>
      {!!outputList?.length && !isLoading && (
        <div className="flex justify-center text-center">
          <Card>
            <CardHeader className="text-center">
              {outputList.map((output, index) => {
                return (
                  <div key={`output_${index}`} className="flex items-center">
                    <CardTitle className="text-lg mr-1">{output?.ingredient}:</CardTitle>
                    <CardDescription className="text-sm pt-0.5">{`${output?.amount} ${output?.unit}`}</CardDescription>
                  </div>
                );
              })}
            </CardHeader>
          </Card>
        </div>
      )}
    </div>
  );
}
