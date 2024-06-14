import axios, { AxiosResponse } from "axios";

let playlistPromise: Promise<AxiosResponse<any>> | null = null;

const postPlaylist = async (url: string) => {
  if (!playlistPromise) {
    playlistPromise = axios.post(
      "http://localhost:8080/playlist",
      { url },
      { withCredentials: true }
    );
  }

  const response = await playlistPromise;
  playlistPromise = null; // Reset the promise so the next call will fetch the data again
  return response.data;
};

export default postPlaylist;
