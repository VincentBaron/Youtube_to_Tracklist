import "./App.css";
import SpotifySignup from "./components/ui/SpotifyLogin";
import SetsPage from "./pages/SetsPage";

function App() {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <div className="mb-7 ">
        <SpotifySignup />
      </div>
      <SetsPage />
    </div>
  );
}

export default App;
