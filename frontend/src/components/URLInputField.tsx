import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useRef } from "react";
import { useCreatePlaylistMutation } from "@/utils/queryOptions";

function URLInputField() {
  const inputRef = useRef<HTMLInputElement>(null);
  const createPlaylistMutation = useCreatePlaylistMutation();

  const handleSubmit = () => {
    const url = inputRef.current?.value;
    if (url) {
      createPlaylistMutation.mutate(url);
    }
  };

  return (
    <div className="flex w-full max-w-sm items-center space-x-2">
      <Input
        ref={inputRef}
        className="flex-grow"
        type="text"
        placeholder="Youtube URL"
      />
      <Button type="button" onClick={handleSubmit}>
        Generate playlist
      </Button>
    </div>
  );
}

export default URLInputField;
