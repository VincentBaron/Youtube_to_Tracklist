import { FC } from "react";
import Homepage from "./pages/HomePage";
import SpotifyLogin from "./components/SpotifyLogin"; // import the SpotifyLogin component

const App: FC = () => {
  return (
    <div>
      <SpotifyLogin /> {/* add the SpotifyLogin component */}
      <Homepage />
    </div>
  );
};

export default App;
