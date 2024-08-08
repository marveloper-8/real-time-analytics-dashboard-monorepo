import Head from "next/head";
import { FC, ReactNode } from "react";

interface LayoutProps {
  children: ReactNode
}

const Layout: FC<LayoutProps> = ({children}) => {
  return (
    <div className="container">
      <Head>
        <title>Real-time Analytics Dashboard</title>
        <meta name="descripton" content="Real-time analytics dashboard using GraphQL and InfluxDB" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>{children}</main>

      <footer>
        <p>&copy; 2024 Your Company Name</p>
      </footer>
    </div>
  )
}