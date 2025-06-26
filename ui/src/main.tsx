import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import "./index.css"
import { Toaster } from "@/components/ui/sonner"
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  useLocation,
} from "react-router"
import { ChatPage, DashboardPage, LoginPage } from "./routes"
import { AuthContextProvider, useAuthContext } from "./context/auth"

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <>
      <Toaster position="bottom-right" closeButton />
      <AuthContextProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Navigate replace to="/dashboard" />} />
            <Route
              path="/chat"
              element={
                <ProtectedRoute>
                  <ChatPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/dashboard"
              element={
                <ProtectedRoute>
                  <DashboardPage />
                </ProtectedRoute>
              }
            />
            <Route path="/login" element={<LoginPage />} />
          </Routes>
        </BrowserRouter>
      </AuthContextProvider>
    </>
  </StrictMode>
)

export function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const { auth, isLoading } = useAuthContext()
  const location = useLocation()

  if (isLoading) {
    return <div>Loading...</div>
  }

  const WANNALOGIN = false // REMOVE
  if (!auth.isLoggedIn || WANNALOGIN) {
    return <Navigate to="/login" state={{ from: location }} replace />
  }

  return children
}
