import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useRef } from "react";

export function InputWithButton() {
  const inputRef = useRef<HTMLInputElement>(null);

  const handleSubmit = async () => {
    console.log(localStorage.getItem("spotifyToken"));
    const url = inputRef.current?.value;

    try {
      const response = await fetch("http://localhost:8080/playlist", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          // Authorization: "Bearer " + localStorage.getItem("spotifyToken"),
        },
        body: JSON.stringify({ url }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }

    // Handle the response data
  };

  return (
    <div className="flex w-full max-w-sm items-center space-x-2">
      <Input
        ref={inputRef}
        className="flex-grow"
        type="email"
        placeholder="Youtube URL"
      />
      <Button type="button" onClick={handleSubmit}>
        Generate playlist
      </Button>
    </div>
  );
}
