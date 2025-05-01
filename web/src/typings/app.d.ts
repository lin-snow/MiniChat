declare namespace App {
  /**
   * Namespace Api
   */
  namespace Api {
    type Response<T> = {
      code: number
      msg: string
      data: T
    }

    namespace Auth {
      type LoginParams = {
        username: string
        password: string
      }

      type LoginResponse = {
        token: string
      }

      type SignupParams = {
        username: string
        password: string
      }
    }

    namespace User {
      type User = {
        id: number
        username: string
        password?: string
        is_admin: boolean
      }

      type UserStatus = {
        user_id: number
        username: string
        is_admin: boolean
      }
    }

    namespace Ech0 {
      type ParamsByPagination = {
        page: number
        pageSize: number
      }

      type Echo = {
        id: number
        content: string
        username: string
        image_url: string
        private: boolean
        user_id: number
        created_at: string
      }

      type EchoToAdd = {
        content: string
        image_url?: string | null
        private: boolean
      }

      type PaginationResult = {
        items: Echo[]
        total: number
      }

      type Status = {
        sys_admin_id: number
        username: string
        users: App.Api.User.UserStatus[]
        total_messages: number
      }

      type HeatMap = {
        date: string
        count: number
      }[]
    }

    namespace Setting {
      type SystemSetting = {
        server_name: string
        allow_register: boolean
      }
    }
  }
}
