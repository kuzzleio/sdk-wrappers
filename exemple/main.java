// /home/kblondel/Downloads/android-studio/jre/bin/javac examples/main.java
// /home/kblondel/Downloads/android-studio/jre/bin/java -Djava.library.path=./ main

public class main {
    public static void main(String argv[]) {
        try {
            System.loadLibrary("kcore");
        } catch (Exception e) {
            e.printStackTrace();
        }

        kuzzle k = kcore.Kuzzle("localhost:7512", "websocket");
    }
}