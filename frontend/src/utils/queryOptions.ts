import { useMutation } from "@tanstack/react-query";
import postPlaylist from "./apiCalls";
import { queryClient } from "../main";

export const useCreatePlaylistMutation = () => {
  return useMutation({
    mutationKey: ["playlist", "create"],
    mutationFn: postPlaylist,
    onSuccess: () => queryClient.invalidateQueries(),
  });
};
