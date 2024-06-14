// import { FC, useState } from "react";
// import { Button } from "@/components/ui/button";
// import axios from "axios";

// const SpotifySignup: FC = () => {
//   const [email, setEmail] = useState("");
//   const [password, setPassword] = useState("");

//   const handleSignup = () => {
//     axios
//       .post("http://localhost:8080/signup", { email, password })
//       .then((response) => {
//         window.location.href = response.data.url;
//         console.log("response data: ", response.data);
//       })
//       .catch((error) => {
//         console.log("error response data:", error.response.data);
//       });
//   };

//   const handleLogin = () => {
//     axios
//       .post(
//         "http://localhost:8080/login",
//         { email, password }
//         // { withCredentials: true }
//       )
//       .then((response) => {
//         console.log("response data: ", response.data);
//       })
//       .catch((error) => {
//         console.log("error response data:", error.response.data);
//       });
//   };

//   return (
//     <div>
//       <input
//         type="email"
//         value={email}
//         onChange={(e) => setEmail(e.target.value)}
//         placeholder="Email"
//       />
//       <input
//         type="password"
//         value={password}
//         onChange={(e) => setPassword(e.target.value)}
//         placeholder="Password"
//       />
//       <Button onClick={handleSignup}>Signup with Spotify</Button>
//       <Button onClick={handleLogin}>Login with Spotify</Button>
//     </div>
//   );
// };

// export default SpotifySignup;

// "use client"

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { toast } from "@/components/ui/use-toast";

const FormSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
});

export function InputForm() {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
    },
  });

  function onSubmit(data: z.infer<typeof FormSchema>) {
    toast({
      title: "You submitted the following values:",
      description: (
        <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
          <code className="text-white">{JSON.stringify(data, null, 2)}</code>
        </pre>
      ),
    });
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="w-2/3 space-y-6">
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input placeholder="shadcn" {...field} />
              </FormControl>
              <FormDescription>
                This is your public display name.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  );
}
