import { FC, useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import axios from "axios";

const SpotifySignup: FC = () => {
  // const [token, setToken] = useState<string | null>(null);

  const handleSignup = () => {
    axios
      .post("http://localhost:8080/signup")
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
      .post("http://localhost:8080/login")
      .then((response) => {
        console.log("response data: ", response.data);
      })
      .catch((error) => {
        console.log("error response data:", error.response.data);
      });
  };

  // useEffect(() => {
  //   fetch("/api/spotify/status", {
  //     method: "GET",
  //   })
  //     .then((response) => response.json())
  //     .then((data) => {
  //       if (data.loggedIn) {
  //         setToken("loggedIn");
  //       }
  //     });
  // }, []);

  // if (token) {
  //   return <div>You are logged in</div>;
  // }

  return (
    <div>
      <Button onClick={handleSignup}>Signup with Spotify</Button>
      <Button onClick={handleLogin}>Login with Spotify</Button>
    </div>
  );
};

export default SpotifySignup;
