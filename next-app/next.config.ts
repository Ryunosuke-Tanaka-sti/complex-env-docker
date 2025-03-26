import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  experimental: {
    optimizePackageImports: ["@chakra-ui/react"],
  }
  // async rewrites() {
  //   return [
  //     {
  //       source: "/api/",
  //       destination: "http://localhost:8080/api/",
  //     },
  //   ];
  // },
};

export default nextConfig;
