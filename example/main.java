import javax.annotation.PostConstruct;
import java.lang.reflect.Method;

public class main {
    public static void main(String argv[]) {
        try {
            kuzzle k = kcore.Kuzzle("localhost:7512", "websocket");

            query_options options = new query_options();
            options.setScroll("yolo");

            try {
                ack_response r = k.createIndex("fgfdgdfg", options);
                System.out.println(r.getAcknowledged());
                System.out.println(r.getShardsAcknowledged());
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
