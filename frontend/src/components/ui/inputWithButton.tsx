import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export function InputWithButton() {
  return (
    <div className="flex w-full max-w-sm items-center space-x-2">
      <Input className="flex-grow" type="email" placeholder="Youtube URL" />
      <Button type="submit">Generate playlist</Button>
    </div>
  );
}
