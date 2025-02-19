import type { Metadata } from "next";
import { Open_Sans } from "next/font/google";
import "./globals.css";

const opensans = Open_Sans({
  variable: "--font-open-sans",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Celebut Admin",
  description: "Re-imagine Your Celebrations with Celebut",
  icons: {
    icon: "/favicon.ico",
    apple: "/apple-touch-icon.png",
  },
  openGraph: {
    title: "Celebut",
    description: "Re-imagine your celebrations with celebut",
    // url: "https://example.com",
    siteName: "Celebut",
    // images: [
    //   {
    //     url: "https://example.com/og-image.png",
    //     width: 1200,
    //     height: 630,
    //     alt: "Website Logo",
    //   },
    // ],
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${opensans.variable} antialiased`}>{children}</body>
    </html>
  );
}
