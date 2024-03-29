import Box from "@mui/joy/Box";
import Typography from "@mui/joy/Typography";
import Stack from "@mui/joy/Stack";
import Link from "@mui/joy/Link";
import { formLabelClasses } from "@mui/joy/FormLabel";
import { useLocation } from "react-router-dom";
import { useEffect } from "react";
import { toast } from "react-toastify";

import { LoginForm } from "../components/LoginForm";

const Login = () => {
  const location = useLocation();
  const from = location.state?.from?.pathname || "/offers";

  useEffect(() => {
    if (
      from !== "/" &&
      from !== "/register" &&
      from !== "/logout" &&
      from !== "/login" &&
      from !== "/offers"
    ) {
      toast.info("You must be logged in to view that page");
    }
  }, [from]);

  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          minHeight: "80dvh",
          width:
            "clamp(var(--Form-maxWidth), (var(--Collapsed-breakpoint) - 100vw) * 999, 100%)",
          maxWidth: "100%",
          px: 2,
        }}
      >
        <Box
          component="main"
          sx={{
            my: "auto",
            py: 2,
            pb: 5,
            display: "flex",
            flexDirection: "column",
            gap: 2,
            width: 400,
            maxWidth: "100%",
            mx: "auto",
            borderRadius: "sm",
            "& form": {
              display: "flex",
              flexDirection: "column",
              gap: 2,
            },
            [`& .${formLabelClasses.asterisk}`]: {
              visibility: "hidden",
            },
          }}
        >
          <Stack gap={4} sx={{ mb: 2 }}>
            <Stack gap={1}>
              <Typography level="h3">Sign in</Typography>
              <Typography level="body-sm">
                New to SkillSnap?{" "}
                <Link href="/register" level="title-sm">
                  Sign up!
                </Link>
              </Typography>
            </Stack>
          </Stack>
          <Stack gap={4} sx={{ mt: 2 }}>
            <LoginForm from={from} />
          </Stack>
        </Box>
      </Box>
    </>
  );
};

export default Login;
