export default class UserState {

    private static _instance:UserState = new UserState();
  
    public state = {};
  
      constructor() {
          if(UserState._instance) {
              throw new Error("Error: Instantiation failed: Use SingletonClass.getInstance() instead of new.");
          }
          UserState._instance = this;
      }
  
      public static getInstance():UserState {
          return UserState._instance;
      }
  }