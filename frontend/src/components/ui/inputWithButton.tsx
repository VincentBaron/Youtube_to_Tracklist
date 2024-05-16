import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useRef } from "react";
import axios from "axios";

export function InputWithButton() {
  const inputRef = useRef<HTMLInputElement>(null);

  const handleSubmit = async () => {
    console.log(localStorage.getItem("spotifyToken"));
    const url = inputRef.current?.value;

    axios
      .post(
        "http://localhost:8080/playlist",
        { url },
        { withCredentials: true }
      )
      .then((response) => {
        console.log("response data: ", response.data);
      })
      .catch((error) => {
        console.log("error response data:", error.response.data);
      });
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
