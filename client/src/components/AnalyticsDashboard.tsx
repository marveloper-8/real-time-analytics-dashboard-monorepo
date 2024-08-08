import { gql, useQuery, useSubscription } from "@apollo/client";
import { FC } from "react";
import Chart from "./Chart";

const GET_ANALYTICS = gql`
  query GetAnalytics($startTime: String!, $endTime: String!) {
    getAnalytics(startTime: $startTime, endTime: $endTime) {
      totalVisits
      uniqueVisitors
      avgSessionDuration
      bounceRate
      topPages {
        url
        visits
      }
    }
  }
`;

const ANALYTICS_SUBSCRIPTION = gql`
  subscription AnalyticsUpdated {
    analyticsUpated {
      totalVisits
      uniqueVisitors
      avgSessionDuration
      bounceRate
      topPages {
        url
        visits
      }
    }
  }
`;

const AnalyticsDashboard: FC = () => {
  const { loading, error, data } = useQuery(GET_ANALYTICS, {
    variables: { startTime: '2024-08-01', endTime: '2024-08-03' },
  });

  useSubscription(ANALYTICS_SUBSCRIPTION, {
    onSubscriptionData: ({ subscriptionData }) => {
      console.log('New analytics data:', subscriptionData.data.analyticsUpated)
    }
  })

  if (loading) return <p>Loading...</p>
  if (error) return <p>Error: {error.message}</p>

  const { totalVisits, uniqueVisitors, avgSessionDuration, bounceRate, topPages } = data.getAnalytics;

  const chartData = {
    labels: topPages.map((page: { url: string }) => page.url),
    datasets: [
      {
        label: 'Page Visits',
        data: topPages.map((page: { visits: number }) => page.visits),
        borderColor: 'rgb(75, 192, 192)',
        backgroundColor: 'rgba(75, 192, 192, 0.5)'
      }
    ]
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8 text-center">Analytics Dashboard</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <div className="bg-white shadow rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-2">Total Visits</h2>
          <p className="text-3xl font-bold">{totalVisits}</p>
        </div>
        <div className="bg-white shadow rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-2">Unique Visitors</h2>
          <p className="text-3xl font-bold">{uniqueVisitors}</p>
        </div>
        <div className="bg-white shadow rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-2">Avg. Session Duration</h2>
          <p className="text-3xl font-bold">{avgSessionDuration.toFixed(2)} seconds</p>
        </div>
        <div className="bg-white shadow rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-2">Bounce Rate</h2>
          <p className="text-3xl font-bold">{(bounceRate * 100).toFixed(2)}%</p>
        </div>
      </div>
      <div className="bg-white shadow rounded-lg p-6">
        <h2 className="text-xl font-semibold mb-4">Top Pages</h2>
        <Chart data={chartData} />
      </div>
    </div>
  )
}

export default AnalyticsDashboard;