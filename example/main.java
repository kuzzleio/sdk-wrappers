import javax.annotation.PostConstruct;
import java.lang.reflect.Method;

public class main {
    public static void main(String argv[]) {
        try {
            Kuzzle k = new Kuzzle("localhost:7512");
            System.out.println(k.checkToken(null));
        } catch(Exception e) {
            e.printStackTrace();
        }
    }
}
