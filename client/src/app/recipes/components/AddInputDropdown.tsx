import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ConversionSystem, ConversionType, InputListProps } from "@/types/recipeTypes";

export default function AddInputDropdown({ inputList, setInputList }: InputListProps) {
  const handleAddInput = (inputSystem: ConversionSystem, type: ConversionType) => {
    setInputList([
      ...inputList,
      {
        ingredient: "",
        inputSystem: inputSystem === "usa" ? "usa" : "metric",
        inputUnit: "",
        outputSystem: inputSystem === "usa" ? "metric" : "usa",
        outputUnit: "",
        type,
        amount: 0,
      },
    ]);
  };
  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger className="m-1 mt-4" asChild>
          <Button>From Metric</Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuItem onClick={() => handleAddInput("metric", "weight")}>
            Weight
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => handleAddInput("metric", "volume")}>
            Volume
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <DropdownMenu>
        <DropdownMenuTrigger className="m-1 mt-4 w-28" asChild>
          <Button>From US</Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuItem onClick={() => handleAddInput("usa", "weight")}>
            Weight
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => handleAddInput("usa", "volume")}>
            Volume
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}
