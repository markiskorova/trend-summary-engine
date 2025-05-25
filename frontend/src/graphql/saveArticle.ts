import { gql } from '@apollo/client'

export const SAVE_ARTICLE_MUTATION = gql`
  mutation SaveArticle($url: String!) {
    saveArticle(url: $url) {
      id
      url
      summary
      createdAt
    }
  }
`
