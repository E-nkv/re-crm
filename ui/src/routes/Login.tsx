import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"

import { toast } from "sonner"
import { useAuthContext, type AuthType } from "@/context/auth"
import { FetchLogin } from "@/api"
import {
  Form,
  FormField,
  FormItem,
  FormMessage,
  FormControl,
  Button,
  FormLabel,
  PasswordInput,
  Input,
} from "@/components/ui"
import { Loader2 } from "lucide-react"
import { Navigate, useLocation, useNavigate } from "react-router"

const loginSchema = z.object({
  nick: z
    .string()
    .min(3, "Nick must be at least 3 characters")
    .nonempty("Nick is required"),
  pass: z
    .string()
    .min(3, "Password must be at least 3 characters")
    .nonempty("Password is required"),
})

type LoginFormValues = z.infer<typeof loginSchema>

export default function LoginPage() {
  const [serverError, setServerError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const nav = useNavigate()
  const location = useLocation()
  const pn = location.state?.from?.pathname
  const from = pn && pn != "/login" ? pn : "/dashboard" //go to whichever route user was trying to go to. if it was originally login, just redirect to dashboard
  const { auth, setAuth } = useAuthContext()

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    mode: "onSubmit",
    defaultValues: {
      nick: "",
      pass: "",
    },
  })

  if (auth.isLoggedIn) {
    console.log("user is logged in. returning to: ", from)
    return <Navigate to={from} replace />
  }
  async function onSubmit(data: LoginFormValues) {
    setServerError(null)
    setLoading(true)

    try {
      const res = await FetchLogin(data.nick, data.pass)
      await new Promise((res) => setTimeout(res, 2500))
      if (!res.ok) {
        if (res.status === 400) {
          setServerError("Invalid username or password")
        } else if (res.status === 404) {
          setServerError("User does not exist")
        } else {
          setServerError("Something went wrong")
        }
        return
      }
      toast("Login successful!")
      const respBody = (await res.json()) as AuthType
      setAuth({ isLoggedIn: respBody.isLoggedIn, role: respBody.role })
      console.log("redirecting to.. ", from)
      nav(from, { replace: true })
    } catch (error) {
      console.log(error)
      setServerError(
        "Network error: Unable to connect to the server. Please check your connection."
      )
    } finally {
      setLoading(false)
    }
  }

  return (
    <Form {...form}>
      <div className="min-h-screen flex items-center justify-center px-4">
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="max-w-md w-full  p-6 bg-color-card dark:bg-gray-900 rounded-lg shadow-md"
          noValidate
        >
          <h2 className="text-2xl font-semibold text-center text-color-card-foreground dark:text-color-card-foreground mb-4 font-stretch-200%">
            Log in, Adventurer!
          </h2>

          {/* Nickname Field */}
          <FormField
            control={form.control}
            name="nick"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-md font-bold">Username *</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Enter your username"
                    {...field}
                    onChange={(e) => {
                      field.onChange(e)
                      if (serverError) setServerError(null)
                    }}
                    disabled={loading}
                    autoComplete="username"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Password Field */}
          <FormField
            control={form.control}
            name="pass"
            render={({ field }) => (
              <FormItem className="mt-3">
                <FormLabel className="text-md font-bold">Password *</FormLabel>
                <FormControl>
                  <PasswordInput
                    placeholder="Enter your password"
                    {...field}
                    name="password"
                    disabled={loading}
                    spellCheck={false}
                    autoComplete="new-password"
                    onChange={(e) => {
                      field.onChange(e)
                      if (serverError) setServerError(null)
                    }}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Server error message */}

          <p
            className={`text-sm text-color-destructive dark:text-destructive mt-1 ${
              serverError == null ? "invisible h-1" : "visible h-1"
            }`}
          >
            {serverError}
          </p>

          {/* Submit Button aligned right */}
          <div className="flex justify-end">
            <Button
              type="submit"
              disabled={loading}
              className="inline-flex items-center space-x-2 px-6"
            >
              {loading ? (
                <div className="w-15 flex justify-center items-center">
                  <Loader2 className="animate-spin size-5" />
                </div>
              ) : (
                <p className="w-15">Log in</p>
              )}
            </Button>
          </div>
        </form>
      </div>
    </Form>
  )
}
