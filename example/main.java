public class main {
    public static void main(String argv[]) {
        try {
            System.loadLibrary("kcore");
        } catch (Exception e) {
            e.printStackTrace();
        }

        try {
            kuzzle k = kcore.Kuzzle("localhost:7512", "websocket");
            System.out.println(k.checkToken("test").getValid());
            System.out.println(k.checkToken("test").getState());
            System.out.println(k.checkToken("test").getExpiresAt());
        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}
