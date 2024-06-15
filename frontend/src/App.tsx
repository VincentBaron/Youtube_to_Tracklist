import { FC } from "react";
// import Homepage from "./pages/HomePage";
import SpotifySignup from "./components/SpotifyLogin"; // import the SpotifyLogin component
import { SetsPlayerPage } from "./pages/SetsPlayerPage";

const App: FC = () => {
  return (
    <div>
      <SpotifySignup /> {/* add the SpotifyLogin component */}
      {/* <Homepage /> */}
      <SetsPlayerPage />
    </div>
  );
};

export default App;
