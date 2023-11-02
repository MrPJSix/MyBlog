export type ArticleReciveData = {
  id: number
  created_at: string
  updated_at: string
  title: string
  content: string
  img: string
  comment_count: number
  read_count: number
  likes: number
  category: {
    category_id: number
    category_name: string
  }
  author: {
    user_id: number
    full_name: string
    avatar_url: string
  }
}
