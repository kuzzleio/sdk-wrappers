import javax.annotation.PostConstruct;
import java.lang.reflect.Method;

public class main {
    public static void main(String argv[]) {
        try {
            kuzzle k = kcore.Kuzzle("localhost:7512", "websocket");

            try {
                token_validity r = k.checkToken(null);
                // System.out.println(r.getState());
            } catch(Exception e) {
                e.printStackTrace();
            }
/*
            System.out.println(k.checkToken("test").getValid());
            System.out.println(k.checkToken("test").getState());
            System.out.println(k.checkToken("test").getExpiresAt());
*/
        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}
