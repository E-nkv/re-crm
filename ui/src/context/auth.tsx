import { FetchMe } from "@/api"
import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  type ReactNode,
} from "react"

type Role = "admin" | "manager" | "sales_rep" | null
export interface AuthType {
  role: Role
  isLoggedIn: boolean
}
interface AuthContextType {
  auth: AuthType
  setAuth: (authType: AuthType) => void
  isLoading: boolean
  setIsLoading: React.Dispatch<React.SetStateAction<boolean>>
}

const RoleContext = createContext<AuthContextType | undefined>(undefined)

interface RoleContextProviderProps {
  children: ReactNode
}

export const AuthContextProvider = ({ children }: RoleContextProviderProps) => {
  const [auth, setAuth] = useState<AuthType>({ isLoggedIn: false, role: null })
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchAuth = async () => {
      try {
        const res = await FetchMe()
        if (!res.ok) {
          setAuth({ role: null, isLoggedIn: false })
          return
        } else {
          const data = (await res.json()) as AuthType
          console.log("user is logged in, with role: ", data.role)
          setAuth({
            role: data.role,
            isLoggedIn: data.isLoggedIn,
          })
        }
      } catch {
        setAuth({ isLoggedIn: false, role: null })
      } finally {
        setIsLoading(false)
      }
    }
    fetchAuth()
  }, [])

  return (
    <RoleContext.Provider value={{ auth, setAuth, isLoading, setIsLoading }}>
      {children}
    </RoleContext.Provider>
  )
}

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = (): AuthContextType => {
  const context = useContext(RoleContext)
  if (!context) {
    throw new Error("useRoleContext must be used within a RoleContextProvider")
  }
  return context
}
