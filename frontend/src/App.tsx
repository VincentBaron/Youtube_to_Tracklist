import { FC } from "react";
import Homepage from "./pages/HomePage";
import SpotifySignup from "./components/SpotifyLogin"; // import the SpotifyLogin component

const App: FC = () => {
  return (
    <div>
      <SpotifySignup /> {/* add the SpotifyLogin component */}
      <Homepage />
    </div>
  );
};

export default App;
