import { ApolloClient, HttpLink, InMemoryCache, split } from "@apollo/client";
import { WebSocketLink } from "@apollo/client/link/ws"
import { getMainDefinition } from "@apollo/client/utilities";

const httpLink = new HttpLink({
  uri: process.env.NEXT_PUBLIC_GRAPHQL_HTTP_URL || 'http://localhost:8080/query',
})

const wsLink = typeof window ? new WebSocketLink({
  uri: process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8080/query',
  options: {
    reconnect: true
  }
}) : null;

const splitLink = typeof window ? split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' && definition.operation === 'subscription'
    );
  },
  wsLink!,
  httpLink
) : httpLink;

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache(),
})

export default client