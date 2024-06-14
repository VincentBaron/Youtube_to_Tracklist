import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/upload")({
  component: Upload,
});

function Upload() {
  return (
    <div className="p-2">
      <h3>Upload a file here</h3>
    </div>
  );
}
