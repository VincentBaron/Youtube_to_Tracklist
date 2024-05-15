import { FC, useEffect, useState } from "react";
import { Button } from "@/components/ui/button";

const SpotifyLogin: FC = () => {
  const [token, setToken] = useState(localStorage.getItem("spotifyToken"));

  const handleLogin = () => {
    const SPOTIFY_AUTHORIZE_ENDPOINT = import.meta.env
      .VITE_SPOTIFY_AUTHORIZE_ENDPOINT;
    const CLIENT_ID = import.meta.env.VITE_CLIENT_ID;
    const REDIRECT_URL_AFTER_LOGIN = import.meta.env
      .VITE_REDIRECT_URL_AFTER_LOGIN;
    const SCOPES_URL_PARAM = import.meta.env.VITE_SCOPES_URL_PARAM;

    window.location.href = `${SPOTIFY_AUTHORIZE_ENDPOINT}?client_id=${CLIENT_ID}&redirect_uri=${REDIRECT_URL_AFTER_LOGIN}&scope=${SCOPES_URL_PARAM}&response_type=token&show_dialog=true`;
  };

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.hash.substring(1));
    const accessToken = urlParams.get("access_token");

    if (accessToken) {
      localStorage.setItem("spotifyToken", accessToken);
      setToken(accessToken);
    }
  }, []);

  if (token) {
    return <div>You are logged in</div>;
  }

  return <Button onClick={handleLogin}>Login with Spotify</Button>;
};

export default SpotifyLogin;
