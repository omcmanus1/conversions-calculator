import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

type selectType = {
  handleChange?: (inp: string, val?: string) => void;
  placeholder: string;
  selectContent: Array<string>;
};

export default function SelectSh({
  handleChange,
  placeholder,
  selectContent,
}: selectType) {
  return (
    <div className="mb-1">
      <Select onValueChange={handleChange}>
        <SelectTrigger>
          <SelectValue placeholder={placeholder} />
        </SelectTrigger>
        <SelectContent>
          {selectContent.map((item) => {
            return (
              <SelectItem key={`${item}_key`} value={item}>
                {item.charAt(0).toUpperCase() + item.slice(1)}
              </SelectItem>
            );
          })}
        </SelectContent>
      </Select>
    </div>
  );
}
