import { FC, useState } from "react";
import { Button } from "@/components/ui/button";
import axios from "axios";

const SpotifySignup: FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSignup = () => {
    axios
      .post("http://localhost:8080/signup", { email, password })
      .then((response) => {
        window.location.href = response.data.url;
        console.log("response data: ", response.data);
      })
      .catch((error) => {
        console.log("error response data:", error.response.data);
      });
  };

  const handleLogin = () => {
    axios
      .post(
        "http://localhost:8080/login",
        { email, password }
        // { withCredentials: true }
      )
      .then((response) => {
        console.log("response data: ", response.data);
      })
      .catch((error) => {
        console.log("error response data:", error.response.data);
      });
  };

  return (
    <div>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
      />
      <Button onClick={handleSignup}>Signup with Spotify</Button>
      <Button onClick={handleLogin}>Login with Spotify</Button>
    </div>
  );
};

export default SpotifySignup;
