import { InputWithButton } from "@/components/ui/inputWithButton";
import { FC } from "react";
import { Card } from "@/components/ui/card";

const App: FC = () => {
  return (
    <div className="flex flex-col items-center pt-10 h-screen">
      <div className="pt-10 mb-4">
        <InputWithButton />
      </div>
      <div className="w-full max-w-6xl">
        <Card>
          <div className="flex">
            <div>
              <img src="/path/to/your/image.jpg" alt="Screenshot" />
            </div>
            <div className="flex-grow flex flex-col">
              <h2 className="text-xl">Title of the YouTube DJ Set</h2>
              <ul>
                <li>
                  <a href="https://open.spotify.com/track/1">Song 1</a>
                </li>
                <li>
                  <a href="https://open.spotify.com/track/2">Song 2</a>
                </li>
                {/* Add more songs as needed */}
              </ul>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default App;
