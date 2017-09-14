import javax.annotation.PostConstruct;
import java.lang.reflect.Method;

public class main {
    public static void main(String argv[]) {
        try {
            kuzzle k = kcore.Kuzzle("localhost:7512", "websocket");

            query_options options = new query_options();
            options.setScroll("yolo");

            try {

                // System.out.println(k.login("local"));
                // ack_response r = k.createMyCredentials("local", options);
                r.getAcknowledged();
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
