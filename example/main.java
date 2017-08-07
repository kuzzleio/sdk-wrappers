public class main {
    public static void main(String argv[]) {
        try {
            System.loadLibrary("kcore");
        } catch (Exception e) {
            e.printStackTrace();
        }

        try {
            kuzzle k = kcore.Kuzzle("localhost:7512", "websocketz");
            System.out.println(k);
        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}